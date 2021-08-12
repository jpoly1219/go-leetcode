#include <iostream>

using namespace std;

void palindrome(const char *s) {
    if (s[0]) {
        cout << s[0];
        if (s[1]) {
            palindrome(s+1);
            cout << s[0];
        }
    }
}

int main() {
    const char *s = "foobar";
    palindrome(s);
    cout << endl;
}