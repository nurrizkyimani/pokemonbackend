
import java.util.ArrayList;
import java.util.List;

class oscilator {
  static void pattern(int n, int compared, List<Integer> rest, boolean positive){

    rest.add(n);

    if(positive == false && n == compared){
      return;
    }

    if(positive == true){
      if(n-5 > 0){
        pattern(n-5, compared, rest, true );
      } else {
        pattern(n-5, compared, rest, false);
      }

    } else {
      pattern(n+5, compared, rest, false);
    }
  
  }

public static void main(String[] args) {
    List<Integer> list = new ArrayList<>();

    int input = 16; 

    pattern(input, input, list, true);

    System.out.println(list);
  }
    
}
