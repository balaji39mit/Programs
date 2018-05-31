#include <bits/stdc++.h>

using namespace std;

string subsequenceAgain(string s, int k) {
    int count[26] = {0};
    for (int i=0;i<s.length();i++)
    {
        int index = s[i] - 'a';
        count[index]++;
    }
    string result = "";
    for (int i=0;i<s.length();i++)
    {
        int index = s[i] - 'a';
        if (count[index] >= k)
        {            
            result += s[i];
        }
    }
    return result;
    // Complete this function
}

int main() {
    string s;
    cin >> s;
    int k;
    cin >> k;
    string result = subsequenceAgain(s, k);
    cout << result << endl;
    return 0;
}

