package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game 实现了ebiten.Game接口
type Game struct {
	x, y  float64 // 方块位置
	speed float64 // 移动速度
	size  float64 // 方块大小
}

// 初始化游戏
func NewGame() *Game {
	return &Game{
		x:     100, // 初始X位置
		y:     100, // 初始Y位置
		speed: 3,   // 移动速度
		size:  50,  // 方块大小
	}
}

// Update 更新游戏状态
func (g *Game) Update() error {
	// 处理键盘输入
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.x -= g.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.x += g.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.y -= g.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.y += g.speed
	}

	// 边界检测
	screenWidth, screenHeight := 640, 480
	if g.x < 0 {
		g.x = 0
	}
	if g.x > float64(screenWidth)-g.size {
		g.x = float64(screenWidth) - g.size
	}
	if g.y < 0 {
		g.y = 0
	}
	if g.y > float64(screenHeight)-g.size {
		g.y = float64(screenHeight) - g.size
	}

	return nil
}

// Draw 渲染游戏画面
func (g *Game) Draw(screen *ebiten.Image) {
	// 填充背景色
	screen.Fill(color.RGBA{30, 30, 30, 255})

	// 绘制方块
	ebitenutil.DrawRect(screen, g.x, g.y, g.size, g.size, color.RGBA{0, 180, 255, 255})
}

// Layout 定义游戏窗口大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // 游戏窗口大小
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("简单移动方块游戏")

	// 启动游戏
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
