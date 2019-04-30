#include <iostream>
using namespace std;

int main() {
  int n, m , r;
  int val;
  int min_s = 1000, max_b = 0;
  cin >> n >> m >> r;
  for (int i = 0; i < n; i++) {
    cin >> val;
    if (min_s >  val) {
      min_s = val;
    }
  } 

  for (int i = 0; i < m; i++) {
    cin >> val;
    if (max_b <  val) {
      max_b = val;
    }
  }
  int maxValue = (r / min_s) * max_b + (r % min_s);
  if (maxValue > r) {
    cout << maxValue;
  }
  else {
    cout << r;
  }
  return 0;
}
