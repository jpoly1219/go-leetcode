import java.util.*;
 
class Palindrome {
    public static void main(String args[]) {
        String original, reverse = "";
        
        original = "abcba";

        int length = original.length();

        for (int i = length - 1; i >= 0; i--)
        reverse = reverse + original.charAt(i);

        if (original.equals(reverse))
            System.out.println("The string is a palindrome.");
        else
            System.out.println("The string isn't a palindrome.");
    }
}