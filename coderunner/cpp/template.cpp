// template file
#include <fstream>
#include <iostream>
#include <string>
#include <stringstream>
#include <vector>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

// insert Solution class here

int main()
{
    Solution sol;

    ifstream i("../testcase-1-two-sum.json");
    json j;
    i >> j;

    vector<vector<int>> vecNums = j["input"]["nums"];
    vector<int> vecTargets = j["input"]["target"];
    vector<vector<int>> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (int i = 0; i < vecNums.size(); i++)
    {
        vector<int> vecSolution = sol.twoSum(vecNums.at(i), vecTargets.at(i));
        if (vecSolution != vecExpected.at(i))
        {
            stringstream input;
            std::copy(vecNums.at(i).begin(), vecNums.at(i).end(), std::ostream_iterator<int>(input, " "));
            input.str();

            stringstream expected;
            std::copy(vecExpected.at(i).begin(), vecExpected.at(i).end(), std::ostream_iterator<int>(expected, " "));
            expected.str();

            stringstream output;
            std::copy(vecSolution.begin(), vecSolution.end(), std::ostream_iterator<int>(output, " "));
            output.str();

            json output = {
                {"result", "wrong"},
                {"input", input.str()},
                {"expected", expected.str()},
                {"output", output.str()}};
            cout << output.dump(4) << endl;
            isOk = false;
            break;
        }
    }
    if (isOk)
    {
        json output = {
            {"result", "OK"}};
        cout << output.dump(4) << endl;
    }
    cout << "test completed" << endl;
    i.close();
}