
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;
import java.util.Deque;
import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.LinkedList;


class fivegraph {

  private boolean adjMatrix[][];
  private int numVertices;
  private int numEdges;

  private ArrayList<Integer>  listdfs = new ArrayList<>();
  private ArrayList<Integer>  listbfs = new ArrayList<>();

  void dfs(int start, boolean[] visited ){
    visited[start] = true;
    listdfs.add(start); 

    for (int i = 0; i < numVertices; i++) { 
      if (adjMatrix[start][i] == true && (!visited[i])) { 
          dfs(i, visited); 
      } 
    } 
  }

  void bfs(int start, boolean[] visited ){
    Queue<Integer> q = new LinkedList<>();
    visited[start] = true;
    q.add(start); 
    
    while (!q.isEmpty()) { 
      listbfs.add(q.peek()); 
      int vis = q.poll();
        // Print the current node 
        
        // For every adjacent vertex to the current vertex 
        for (int i = 0; i < numVertices; i++) { 
            if (adjMatrix[vis][i] == true && (!visited[i])) { 
                q.add(i); 
                
                visited[i] = true; 

            } 
        } 
    } 
  }

  public fivegraph(int numVertices, int numEdges) {
    this.numVertices = numVertices;
    this.numEdges = numEdges;
    adjMatrix = new boolean[numVertices][numVertices];
  }

  public void addEdge(int i, int j) {
    adjMatrix[i][j] = true;
    adjMatrix[j][i] = true;
  }
  

public static void main(String[] args) {

    //create graph using adjecency matrix;
    int v = 5;
    int w = 4;
    fivegraph g = new fivegraph(v, w);

    //add the relationship;
    g.addEdge(0,1);
    g.addEdge(0,2);
    g.addEdge(0,3);
    g.addEdge(1,4);

    //run DFS
    System.out.println("DFS ");
    boolean[] visited = new boolean[v];
    Arrays.fill(visited, false);
    g.dfs(0, visited);
    System.out.println(g.listdfs);

    //RUN BFS
    System.out.println("BFS ");
    boolean[] visitedbfs = new boolean[v];
    g.bfs(0, visitedbfs);
    System.out.println(g.listbfs);
    
  }
    
}
