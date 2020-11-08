
import java.util.Arrays;
import java.util.HashMap;
import java.util.ArrayList;


class smallestElement {
  static void helper(){
      
    
  }

public static void main(String[] args) {
    HashMap<Integer, Integer> map = new HashMap<>();
    
    int[] inp = new int[]{3, 5, 6, 2, 3 ,9 ,12 ,4 ,3 ,7};

    int[] nums = new int[]{3, 5, 6, 2, 3 ,9 ,12 ,4 ,3 ,7};

    Arrays.sort(nums);

    HashMap<Integer, Integer> map2 = new HashMap<>();

    int index = 1; 
    int indexmap = 0;
    for( int num : nums){
        if(!map.containsKey(num)){
          map.put(num, index++);
          map2.put(indexmap++, num); 
        }
    }

    int[] out = new int[inp.length];

    for(int i=0; i < inp.length ; i++ ){
      // System.out.println(map2.get(map.get(inp[i])));
      if(map2.get(map.get(inp[i])) == null){
        out[i] =  99 ;
        continue;
      }

      out[i] = map2.get(map.get(inp[i]));
    }
    System.out.println(Arrays.toString(out));
  }
    
}
