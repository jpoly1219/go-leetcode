import java.io.*;
import java.util.*;
 
class GFG
{
 
    static String str;
 
    // Utility function to calculate
    // minimum of two numbers
    static int min(int a, int b)
    {
        return (a < b) ? a : b;
    }
 
    // Function to calculate maximum
    // overlap in two given strings
    static int findOverlappingPair(String str1,
                                   String str2)
    {
         
        // max will store maximum
        // overlap i.e maximum
        // length of the matching
        // prefix and suffix
        int max = Integer.MIN_VALUE;
        int len1 = str1.length();
        int len2 = str2.length();
 
        // check suffix of str1 matches
        // with prefix of str2
        for (int i = 1; i <=
                            min(len1, len2); i++)
        {
 
            // compare last i characters
            // in str1 with first i
            // characters in str2
            if (str1.substring(len1 - i).compareTo(
                        str2.substring(0, i)) == 0)
            {
                if (max < i)
                {
 
                    // Update max and str
                    max = i;
                    str = str1 + str2.substring(i);
                }
            }
        }
 
        // check prefix of str1 matches
        // with suffix of str2
        for (int i = 1; i <=
                           min(len1, len2); i++)
        {
 
            // compare first i characters
            // in str1 with last i
            // characters in str2
            if (str1.substring(0, i).compareTo(
                      str2.substring(len2 - i)) == 0)
            {
                if (max < i)
                {
 
                    // pdate max and str
                    max = i;
                    str = str2 + str1.substring(i);
                }
            }
        }
 
        return max;
    }
 
    // Function to calculate smallest
    // string that contains
    // each string in the given set as substring.
    static String findShortestSuperstring(
                          String arr[], int len)
    {
         
        // run len-1 times to consider every pair
        while (len != 1)
        {
             
            // To store maximum overlap
            int max = Integer.MIN_VALUE;
           
            // To store array index of strings
            // involved in maximum overlap
            int l = 0, r = 0;
                  
            // to store resultant string after
            // maximum overlap
            String resStr = "";
 
            for (int i = 0; i < len; i++)
            {
                for (int j = i + 1; j < len; j++)
                {
 
                    // res will store maximum
                    // length of the matching
                    // prefix and suffix str is
                    // passed by reference and
                    // will store the resultant
                    // string after maximum
                    // overlap of arr[i] and arr[j],
                    // if any.
                    int res = findOverlappingPair
                                  (arr[i], arr[j]);
 
                    // Check for maximum overlap
                    if (max < res)
                    {
                        max = res;
                        resStr = str;
                        l = i;
                        r = j;
                    }
                }
            }
 
            // Ignore last element in next cycle
            len--;
 
            // If no overlap,
            // append arr[len] to arr[0]
            if (max == Integer.MIN_VALUE)
                arr[0] += arr[len];
            else
            {
               
                // Copy resultant string
                // to index l
                arr[l] = resStr;
               
                // Copy string at last index
                // to index r
                arr[r] = arr[len];
            }
        }
        return arr[0];
    }
 
    // Driver Code
    public static void main(String[] args)
    {
        String[] arr = { "catgc", "ctaagt",
                      "gcta", "ttca", "atgcatc" };
        int len = arr.length;
 
        System.out.println("The Shortest Superstring is " +
                        findShortestSuperstring(arr, len));
    }
}