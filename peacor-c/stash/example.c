// from: http://blog.chrislowis.co.uk/2008/11/24/ruby-gsl-pearson.html
#include <math.h>
double inline_pearson(int n, VALUE x, VALUE y) {
  double sum1 = 0.0;
  double sum2 = 0.0;
  double sum1Sq = 0.0;
  double sum2Sq = 0.0;
  double pSum = 0.0;
  
  VALUE *x_a = RARRAY(x)->ptr;
  VALUE *y_a = RARRAY(y)->ptr;
  
  int i;
  for(i=0; i<n; i++) {
    double this_x;
    double this_y;
    this_x = NUM2DBL(x_a[i]);
    this_y = NUM2DBL(y_a[i]);
    sum1 += this_x;
    sum2 += this_y;
    sum1Sq += pow(this_x, 2);
    sum2Sq += pow(this_y, 2);
    pSum += this_y * this_x;
  }
  double num;
  double den;
  num = pSum - ( ( sum1 * sum2 ) / n );
  den = sqrt( ( sum1Sq - ( pow(sum1, 2) ) / n ) *
        ( sum2Sq - ( pow(sum2, 2) ) / n ) );
  if(den == 0){
    return 0.0;
  } else {
    return num / den;
  }
}
