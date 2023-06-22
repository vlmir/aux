
int ParseTab
(
FILE *fpr, // file pointer for reading
double **data, // for data values
char **meta, // for labels
int n, // number of rows
int  m, // number of data columns
int STRMX, // max lenght of symbols
int DBLMX // max length of data values
)

{
	int i = 0;
	char *dlm = "\t";
  while (!feof(fpr))
  {
		char label[STRMX];
		char vals[m*DBLMX];
		if ((fscanf(fpr,"%s\t%[\t-.1234567890]\n", label, vals)) == 2)
		{
			strcpy (meta[i], label); // SIC!
  		char *field, *string;
  		string = strdup(vals);
  		int j = 0;
      while((field = strsep(&string, dlm)) != NULL )
  		{
  			double val = atof (field);
  			data[i][j] = val;
      	j++;
  		}
  		i++;
			if (j != m)
			{
		 		char *msg = "Unexpected number of columns";
				fprintf(stderr, "%s: %s: %d in row: %d \n", __func__, msg, j, i);
			}
		}
	}
	if (i != n) 
	{
 		char *msg = "Unexpected number of rows";
 		fprintf(stderr, "%s: %s: %d \n", __func__, msg, i);
	}
	return (0);
}
