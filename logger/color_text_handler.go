package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"
	"time"
)

type ColorTextHandler struct {
	h slog.Handler
	b *bytes.Buffer
	m *sync.Mutex
}

func NewColorTextHandler(opts *slog.HandlerOptions) *ColorTextHandler {
	b := &bytes.Buffer{}

	if opts == nil {
		opts = &slog.HandlerOptions{}
	}

	return &ColorTextHandler{
		b: b,
		m: &sync.Mutex{},
		h: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: suppressDefaults(opts.ReplaceAttr),
		}),
	}
}

func (self *ColorTextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return self.h.Enabled(ctx, level)
}

func (self *ColorTextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	self.h = self.h.WithAttrs(attrs)
	return self
}

func (self *ColorTextHandler) WithGroup(name string) slog.Handler {
	self.h = self.h.WithGroup(name)
	return self
}

func (self *ColorTextHandler) Handle(ctx context.Context, r slog.Record) error {
	lvl := Text("[" + r.Level.String() + "]")
	name := Text("")
	attrs, err := self._attrs(ctx, r)

	if err != nil {
		return err
	}

	if v, exists := attrs["name"].(string); exists {
		name = Text(v)
	} else {
		return errors.New("`name` attribute is required")
	}

	if !Match(name.String(), os.Getenv("LOG")) {
		return nil
	}

	switch r.Level {
	case slog.LevelDebug:
		lvl = lvl.BlueForeground().Bold()
		name = name.BlueForeground().Bold()
	case slog.LevelInfo:
		lvl = lvl.CyanForeground().Bold()
		name = name.CyanForeground().Bold()
	case slog.LevelWarn:
		lvl = lvl.YellowForeground().Bold()
		name = name.YellowForeground().Bold()
	case slog.LevelError:
		lvl = lvl.RedForeground().Bold()
		name = name.RedForeground().Bold()
	}

	delete(attrs, "name")
	b, err := json.MarshalIndent(attrs, "", "  ")

	if err != nil {
		return err
	}

	lines := strings.Split(r.Message, "\n")

	for _, line := range lines {
		fmt.Println(
			Text(r.Time.Format(time.RFC3339)).GrayForeground(),
			lvl.String(),
			name.String(),
			line,
		)
	}

	if r.Level >= slog.LevelWarn {
		fmt.Println(Text(b).GrayForeground())
	}

	return nil
}

func (self *ColorTextHandler) _attrs(ctx context.Context, r slog.Record) (map[string]any, error) {
	self.m.Lock()

	defer func() {
		self.b.Reset()
		self.m.Unlock()
	}()

	if err := self.h.Handle(ctx, r); err != nil {
		return nil, err
	}

	attrs := map[string]any{}

	if err := json.Unmarshal(self.b.Bytes(), &attrs); err != nil {
		return nil, err
	}

	return attrs, nil
}

func suppressDefaults(next func([]string, slog.Attr) slog.Attr) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey || a.Key == slog.LevelKey || a.Key == slog.MessageKey {
			return slog.Attr{}
		}

		if next == nil {
			return a
		}

		return next(groups, a)
	}
}
