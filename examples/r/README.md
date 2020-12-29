# Example Jupyter Notebooks to demo the use of `ncbi-datasets` R library

These Jupyter Notebooks demonstrate how you can use the `ncbi-datasets` R  library to query and download data using NCBI Datasets.  They can also be run on binder as an RStudio markdown file.

## Usage

Click on the links below to launch a non-interactive viewer to explore how to use the library. For an interactive session, we recommend using Binder, either as a Jupyter notebook, or RStudio.


| File | Non-interactive | Jupyter Notebook | RStudio |
| --- | --- | --- | --- |
| ncbi_datasets_gene_notebook.ipynb | [ncbi_datasets_gene_notebook.ipynb](ncbi_datasets_gene_notebook.ipynb) | [![Binder](https://mybinder.org/badge_logo.svg)](https://mybinder.org/v2/gh/ncbi/datasets/master?filepath=examples/r/ncbi_datasets_gene_notebook.ipynb) | [![Binder](https://mybinder.org/badge_logo.svg)](https://mybinder.org/v2/gh/ncbi/datasets/master?filepath=examples/r/ncbi_datasets_r_notebook.Rmd&urlpath=rstudio)

## Setup local Jupyter server

```bash
python -m venv ve_test
source ve_test/bin/activate
pip install jupyter
Rscript - <<EOF
install.packages('IRkernel', repos='http://cran.us.r-project.org')
IRkernel::installspec(user = TRUE)
EOF

jupyter-notebook --no-browser --port=<port number> --ip=0.0.0.0
```
