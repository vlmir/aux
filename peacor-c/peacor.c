#include "peacor.h"
/*###########################################################################*/
/*
 input data file:
 tab-delimited
 rows: variables
 columns: observations
 variable's symbols in the first column 
 all the other columns contain variable's values
*/
int Parse
(
FILE *fpr, // file pointer for reading
double **data, // for data values
char **meta, // for labels
int n, // number of rows
int  m, // number of data columns
char *dlm // field delimiter
)

{
	const int LNMX = STRMX + m*DBLMX;
	char line[LNMX];
	int i = 0;
	while (fgets (line, LNMX, fpr))
	{
		char *field;
		field = strtok (line, dlm);
  	int j =  -1;
		while (field != NULL)
		{
			if (j == -1)
			{
				strcpy (meta[i], field); // SIC!
			}
			else
			{
  			double val = atof (field);
  			data[i][j] = val;
  		}
			j++;
			field = strtok (NULL, dlm);
		}
		if (j != m)
		{
	 		char *msg = "Unexpected number of columns";
			fprintf(stderr, "%s: %s: %d in row: %d \n", __func__, msg, j, i);
		}
		i++;
	}
	if (i != n) 
	{
 		char *msg = "Unexpected number of rows";
 		fprintf(stderr, "%s: %s: %d \n", __func__, msg, i);
	}
	return (0);
}

/*###########################################################################*/
int main ( int argc, char *argv[] )
{
	const char *help[7] = {
	"-a: export absolute values",
	"-f: export full matrix, default: no main diagonal",
	"-h: help",
	"-d: delimiter, default: tab or space",
	"-r: number of rows",
	"-c: number of columns",
	"-i: input data file",
	};
	
  extern char *optarg;
  int opt;
	char *opts = "afhvd:i:r:c:";
	char *infn; // path to the data file
	int n; // number of variates
	int m; // number of observations
	char *dlm = "\t "; // default field delimiter (white space)
	/// flags
  int ierr = 0;
	int aa = 0; // '1': exporting absolute values
	int ff = 0; // '0': no main diagonal
	int vv = 0; // '1': print version
  while ((opt = getopt(argc, argv, opts)) != EOF)
    switch (opt) {
      case 'i':
        infn = optarg;
        break;
      case 'd':
        dlm = optarg;
        break;
      case 'r':
        n = atoi(optarg);
        break;
      case 'c':
        m = atoi(optarg);
        break;
      case 'v':
        vv = 1;
        break;
      case 'f':
        ff = 1;
        break;
      case 'a':
        aa = 1;
        break;
      case 'h':
        ierr = 1;
        break;
      case '?':
        ierr = 1;
      break;
    }
  if (infn == NULL)
    ierr = 1;
  if (ierr)
	{
		usage(argv[0], help, 7);
		exit(1);
	}
	if (vv)
	{
		fprintf(stderr, "Version: %.2f \n", VERSION);
		exit(1);
	}
  FILE *fpr; // for reading
  if ((fpr = fopen(infn, "r")) == NULL)
	{
		fprintf(stderr, "%s:fpr: failed to open file: %s \n", __func__, infn);
		exit(1);
	}

  double **data;
  CallocD2d(&data, n, m);
  char **meta;
  CallocC2d(&meta, n, STRMX);

	Parse (fpr, data, meta, n, m, dlm);

	for (int i=0; i<n; i++)
	{
		double v[m];
		for (int k=0; k<m; k++) {v[k] = data[i][k];}
		for (int j=0; j<n; j++)
		{
			if (ff == 0 && j == i) continue;
			double u[m];
			for (int k=0; k<m; k++) {u[k] = data[j][k];}
			double pcc = gsl_stats_correlation (v, 1, u, 1, m);
			if (aa) pcc = fabs(pcc);
			printf ("%s\t%s\t%1.3f\n", meta[i], meta[j], pcc);
		}
	}

	FreeD2d (&data);
	FreeC2d (&meta);
  return 0;
}
