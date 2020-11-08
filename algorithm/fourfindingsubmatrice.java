
import java.util.Arrays;
import java.util.HashMap;
import java.util.ArrayList;


class fourfindingsubmatrice {
  static int[] helper(int[][] matrix, int[][] matrixtofind){


    int kn = matrixtofind.length;
    int ln = matrixtofind[0].length;

    System.out.println(kn);
    System.out.println(ln);
      
    for(int i=0; i < matrix.length; i++){
      for(int j=0; j < matrix[0].length; j++){
        boolean isComplete = false; 
        //if matrix start is equal with matrid find;
        if(matrix[i][j] == matrixtofind[0][0]){
            int o = i;
            int p = j;
            for(int k= 0; k < kn; k++){
              p = j;   
              for(int l= 0; l < ln ; l++){
                System.out.println("===========");
                if(matrix[o][p] == matrixtofind[k][l] ){
                  isComplete = true; 
                  p++;
                
                } else  {
                  // if unmatched
                  // return new int[]{0,0};
                  isComplete = false; 
                }    
              }
              o++;
            }
            if(isComplete){
                return new int[]{i,j};
            } 
        }
      }
    }
    return new int[]{9,9};
    
  }

public static void main(String[] args) {
  // int n = 5;
  // int m = 5;   
  
  int[][] matrix = new int[][]{
    {1, 2, 3, 4, 5},
    {6 ,7 ,8 ,9 ,1},
    {2 ,3 ,4 ,5 ,6},
    {7, 8, 9, 1, 2},
    {3, 4, 5, 6, 7}};

  int[][] matrixtofind = new int[][]{
    {7, 8, 9},
    {3, 4, 5},
    {8, 9, 1}
  };
  
  int[] test = helper(matrix, matrixtofind);
  System.out.println(Arrays.toString(test));

  }
}
