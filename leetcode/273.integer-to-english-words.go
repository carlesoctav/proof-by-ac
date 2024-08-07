package main

import "strconv"

// @leet start
func numberToWords(num int) string {
    modifier := make(map[int]string)
    modifier[9] = "Billion"
    modifier[6] = "Million"
    modifier[3] = "Thousand"
    modifier[2] = "Hundred"
    modifier[1] = "Ten"
    mapping_digit := map[string]string{
        "1":  "One",
        "2":  "Two",
        "3":  "Three", // Keeping it as per your original definition
        "4":  "Four",
        "5":  "Five",
        "6":  "Six",
        "7":  "Seven",
        "8":  "Eight",
        "9":  "Nine",
        "10": "Ten",
    }
    ten_modifier := map[string]string{
        "11": "Eleven",
        "12": "Twelve",
        "13": "Thirteen",
        "14": "Fourteen",
        "15": "Fifteen",
        "16": "Sixteen",
        "17": "Seventeen",
        "18": "Eighteen",
        "19": "Nineteen",
        "2":  "Twenty",
        "3":  "Thirty",
        "4":  "Forty",
        "5":  "Fifty",
        "6":  "Sixty",
        "7":  "Seventy",
        "8":  "Eighty",
        "9":  "Ninety",
    }

    ans := ""
    num_str := strconv.Itoa(num) 
    i := 0
    for i < len(num_str){
        take_modifier := len(num_str) - i - 1
        
        if string( num_str[i] ) == "0" {
            i++
            continue
        }

        if take_modifier % 3 == 0 {
            mod := take_modifier / 3
            ans += mapping_digit[string( num_str[i] )] + " "+ modifier[mod]+ " "
        } else if take_modifier % 3 == 2 {
            mod := take_modifier 
            ans += mapping_digit[string( num_str[i] )] + " "+ modifier[mod]+ " "
        } else if take_modifier % 3 == 1 {
            if string(num_str[i]) == "1"{
                to_append := string(num_str[i])
                i++
                to_append += string(num_str[i])
                ans += ten_modifier[to_append]+ " "
            } else {
                ans += ten_modifier[string( num_str[i] )]+ " "
            }
        }
        i++;
    }

    return ans
    
}
// @leet end
