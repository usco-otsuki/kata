#include <iostream>
#include <stdio.h>
#include <string>

using namespace std;

string m[2000] = {""};
int a_up[2000][2000] = {{0}};
int a_left[2000][2000] = {{0}};
int a_down[2000][2000] = {{0}};
int a_right[2000][2000] = {{0}};
int total[2000][2000] = {{0}};


int main() {
  int H, W;

  cin >> H >> W;
  string str;
  for (int i = 0; i < H; i++) {
    cin >> str;
    m[i] = str;
  } 
  for (int h = 0; h < H; h++) {
    for(int w = 0; w < W; w++) {
      if (m[h][w] == '#') {
        a_left[h][w] = 0;
        a_up[h][w] = 0;
        continue;
      }
      a_left[h][w] = 1;
      if (w > 0) {
        a_left[h][w] += a_left[h][w-1];
      }
      a_up[h][w] = 1;
        if (h > 0){
          a_up[h][w] += a_up[h-1][w];
        }
    }
  }

  for (int h = H-1; h >= 0; h--) {
    for(int w = W-1; w >= 0; w--) {
      if (m[h][w] == '#') {
        a_right[h][w] = 0;
        a_down[h][w] = 0;
        continue;
      }
      a_right[h][w] = 1;
      if (w < W - 1) {
        a_right[h][w] += a_right[h][w+1];
      }
      a_down[h][w] = 1;
        if (h < H - 1){
          a_down[h][w] += a_down[h+1][w];
        }
    }
  }

  int maxVal = 0;
  for (int h = 0; h < H; h++) {
    for(int w = 0; w < W; w++) {
      total[h][w] = a_left[h][w] + a_right[h][w] + a_up[h][w] + a_down[h][w];
      total[h][w] -= 3;
      if (maxVal < total[h][w]) {
        maxVal = total[h][w];
      }
    }
  }
  cout << maxVal << endl;

  return 0;
}
