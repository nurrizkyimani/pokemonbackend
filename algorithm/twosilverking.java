
import java.util.ArrayList;
import java.util.List;

class twosilverking {

  public static int helper(int input, int output ){

    if(input / 5 ==  0){
      return output;
    }

    int res = input / 5;
    // System.out.println("this is  res " + res);
    
    int outputt =  helper(res, output += res);

    // System.out.println(output);

    return outputt;

  }

 

  public static void main(String[] args) {
    int input = 30;
    int output = input; 
    
    int ouput2 = helper(30, output);

    System.out.print(ouput2);


  }
    
}
