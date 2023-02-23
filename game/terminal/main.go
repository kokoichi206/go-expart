package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	NOTHING = 0
	WALL    = 1
	PLAYER  = 69

	MAX_SAMPLES = 100
)

type input struct {
	pressedKey byte
}

func (i *input) update() {
	// only 1 byte
	b := make([]byte, 1)
	os.Stdin.Read(b)

	i.pressedKey = b[0]
}

type position struct {
	x int
	y int
}

type player struct {
	pos   position
	level *level
	input *input

	reverse bool
}

func (p *player) update() {
	if p.reverse {
		p.pos.x -= 1

		if p.pos.x == 2 {
			p.reverse = false
		}
		return
	}

	p.pos.x += 1
	if p.pos.x == p.level.width-2 {
		p.reverse = true
	}
}

func (l *level) set(pos position, v int) {
	l.data[pos.y][pos.x] = v
}

type level struct {
	width  int
	height int
	data   [][]int
}

func newLevel(width, height int) *level {
	data := make([][]int, height)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]int, width)
		}
	}

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if w == 0 || w == width-1 {
				data[h][w] = WALL
			}
			if h == 0 || h == height-1 {
				data[h][w] = WALL
			}
		}
	}
	return &level{
		width:  width,
		height: height,
		data:   data,
	}
}

type game struct {
	isRunning bool
	level     *level
	stats     *stats

	player *player
	input  *input

	drawBuf *bytes.Buffer
}

func newGame(width, height int) *game {
	// enter が押されるまで入力を待たない！
	// m1 mac で機能してないかも, ラズパイでは確認済み
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// 入力された文字を打たない！
	// m1 mac で機能してないかも, ラズパイでは確認済み
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var (
		lvl   = newLevel(width, height)
		input = &input{}
	)

	return &game{
		level:   lvl,
		drawBuf: new(bytes.Buffer),
		stats:   newStats(),

		player: &player{
			level: lvl,
			pos: position{
				x: 2,
				y: 5,
			},
			input: input,
		},
		input: input,
	}
}

func (g *game) start() {
	g.isRunning = true
	g.loop()
}

func (g *game) loop() {
	for g.isRunning {
		// start := time.Now()
		g.input.update()

		g.update()
		g.render()
		g.stats.update()
		// delta := time.Since(start).Milliseconds()
		// diff
		// limit fps!
		time.Sleep(time.Millisecond * 15)
	}
}

func (g *game) update() {
	// listen inputs
	// blocking!!
	// os.Stdin.SetReadDeadline(time.Now().Add(time.Minute))
	// os.Stdin(b)
	g.level.set(g.player.pos, NOTHING)
	g.player.update()
	g.level.set(g.player.pos, PLAYER)
}

type stats struct {
	start  time.Time
	frames int
	fps    float64
}

func newStats() *stats {
	return &stats{
		fps:   69,
		start: time.Now(),
	}
}

func (s *stats) update() {
	s.frames++

	if s.frames == MAX_SAMPLES {
		s.fps = float64(s.frames) / time.Since(s.start).Seconds()
		s.frames = 0
		s.start = time.Now()
	}
}

func (g *game) render() {
	g.drawBuf.Reset()
	fmt.Fprint(os.Stdout, "\033[2J\033[1;1H")

	g.renderLevel()
	g.renderStats()
	fmt.Fprint(os.Stdout, g.drawBuf.String())
}

func (g *game) renderStats() {
	g.drawBuf.WriteString("-- STATS\n")
	g.drawBuf.WriteString(fmt.Sprintf("fps: %.2f\n", g.stats.fps))
}

func (g *game) renderLevel() {
	// unicode characters
	// https://www.w3.org/TR/xml-entity-names/025.html
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == NOTHING {
				g.drawBuf.WriteString(" ")
			}
			if g.level.data[h][w] == WALL {
				g.drawBuf.WriteString("▒")
			}
			if g.level.data[h][w] == PLAYER {
				g.drawBuf.WriteString("☭")
			}
		}
		g.drawBuf.WriteString("\n")
	}
}

func main() {
	width := 80
	height := 18

	g := newGame(width, height)
	g.start()
}
