# Similarity Hashing

The MinHash signature and similarities are calculated using the minhash golang library.
The initial parts of sharding (slicing the input into pairs of characters) are done in the utilis.go file.

Other files include functionalities such as:
* construct example hostnames and flag sequences
* have functionalities to perturb these
* plot the similarity between perturbed strings (this is done using Python and setting up a python environment may be necessary)
* compare how accurately LSH approximates the Jacardi Similarity ( i.e. size of set intersection / size of set unition)

All the experiments are run from the main.go file.

To install the required packages

```bash
    pip install -r requirements.txt
```
