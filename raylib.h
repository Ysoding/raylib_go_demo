
#include <stdbool.h>
// typedef enum bool { false = 0, true = !false } bool;

// Color, 4 components, R8G8B8A8 (32bit)
typedef struct Color {
  unsigned char r; // Color red value
  unsigned char g; // Color green value
  unsigned char b; // Color blue value
  unsigned char a; // Color alpha value
} Color;

// Vector2, 2 components
typedef struct Vector2 {
  float x; // Vector x component
  float y; // Vector y component
} Vector2;

extern void
InitWindow(int width, int height,
           const char *title); // Initialize window and OpenGL context

extern bool
WindowShouldClose(void);       // Check if application should close (KEY_ESCAPE
                               // pressed or windows close icon clicked)
extern void CloseWindow(void); // Close window and unload OpenGL context

extern void SetTargetFPS(int fps); // Set target FPS (maximum)

extern float
GetFrameTime(void); // Get time in seconds for last frame drawn (delta time)

extern void
ClearBackground(Color color); // Set background color (framebuffer clear color)
extern void BeginDrawing(void); // Setup canvas (framebuffer) to start drawing
extern void
EndDrawing(void); // End canvas drawing and swap buffers (double buffering)

extern void
DrawRectangleV(Vector2 position, Vector2 size,
               Color color); // Draw a color-filled rectangle (Vector version)

extern int GetScreenWidth(void);  // Get current screen width
extern int GetScreenHeight(void); // Get current screen height

extern Color ColorFromHSV(float hue, float saturation,
                          float value); // Get a Color from HSV values, hue
                                        // [0..360], saturation/value [0..1]

extern int GetRandomValue(
    int min, int max); // Get a random value between min and max (both included)