files=(scripts/rebase_patches/*)
nfiles=${#files[@]}
for ((i = 0; i < nfiles; i += 5)); do
    for ((j = i; j < nfiles; j++)); do
        git apply ${files[$j]}
    done
done
