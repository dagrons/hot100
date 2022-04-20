package com.company;

import java.util.Scanner;

public class Main3 {
    static int[] states = new int[10];
    public static void main(String[] args) {
		int[] states =  int[]{
			0b1111110,
			0b0110000,
			0b1101101,
			0b1111001,
			0b0110011,
			0b1011011,
			0b1011111,
			0b1110000,
			0b1111111,
			0b1111011
		};
        Scanner scanner = new Scanner(System.in);
        String str = scanner.nextLine();

        int res= 0;
        int start = 0;
        for(char ch:str.toCharArray() ){
            res += countOne(states[(ch-'0')]^start);
            start = states[(ch-'0')];
        }
        System.out.println(res);
    }

    public static int countOne(int t){
        int count = 0;
        while(t>0){
            t=t-(t&(-t));
            
            count++;
        }
        return count;
    }
}