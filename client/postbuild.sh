echo "# Running the post build script"
echo "# Removing the old dist folder if any"
rm -rf ../dist

echo "# Copying the newly created dist folder to the root folder"
mv dist ../