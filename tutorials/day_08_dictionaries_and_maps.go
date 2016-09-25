package tutorials
/*
#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include <map>
#include <string>
using namespace std;


int main() { 
    int n = 0;
    scanf("%d", &n);
    map<string,long long int> np;
    char name[64] = {0};
    long long int phone;
    for (int i = 0; i < n; ++i) {
        scanf("%s%lld", name, &phone);
        np.insert(std::pair<string, int>(string(name), phone));
    }
    cin.ignore();
    while(cin >> name ) {
        auto it = np.find(string(name));
        if (it != np.end() ) {
            printf("%s=%lld\n", name, it->second);
        } 
        else {
            printf("Not found\n");
        }
    }
    return 0;
}
*/