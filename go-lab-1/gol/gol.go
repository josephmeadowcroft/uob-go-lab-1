package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	var sum int = 0
	var IMHT, IMWD int = p.imageHeight, p.imageWidth

	for y = 0; y < IMHT; y++ {
    for x = 0; x < IMWD; x++ {
      sum = A[(y+IMHT-1)%IMHT][(x+IMWD-1)%IMWD] + A[(y+IMHT-1)%IMHT][( x+IMWD )%IMWD] + A[(y+IMHT-1)%IMHT][(x+IMWD+1)%IMWD] + \
            A[( y+IMHT )%IMHT][(x+IMWD-1)%IMWD]                    +                    A[( y+IMHT )%IMHT][(x+IMWD+1)%IMWD] + \
            A[(y+IMHT+1)%IMHT][(x+IMWD-1)%IMWD] + A[(y+IMHT+1)%IMHT][( x+IMWD )%IMWD] + A[(y+IMHT+1)%IMHT][(x+IMWD+1)%IMWD];
      if(A[y][x] == 1)
      {
        if (sum < 2)
           B[y][x] = 0;
        else if (sum == 2 || sum == 3)
           B[y][x] = 1;
        else
           B[y][x] = 0;
      }
      else
      {
        if (sum == 3)
          B[y][x] = 1;
        else
          B[y][x] = 0;
      }
    }
  }

	return world
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
