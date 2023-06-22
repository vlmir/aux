#include <stdio.h>
#include <gsl/gsl_rng.h>
#include <gsl/gsl_statistics.h>
#include <gsl/gsl_statistics_int.h>

gsl_rng * r;  /* global generator */

int
main (void)
{
  const gsl_rng_type * T;

  gsl_rng_env_setup();

  T = gsl_rng_default;
  r = gsl_rng_alloc (T);

  printf ("# generator type: %s\n", gsl_rng_name (r));
  printf ("# seed = %lu\n", gsl_rng_default_seed);

	int n = 3;
	int m = 5;
	for (int i=0; i<n; i++)
	{
		int v[m];
		for (int j=0; j<m; j++)
			{ v[j] = gsl_rng_get (r); }
		double mean = gsl_stats_int_mean (v, 1, m);
		for (int j=0; j<m; j++)
			{printf ("%f ", v[j]/mean);}
		printf ("\n");
	}

  gsl_rng_free (r);
  return 0;
}
