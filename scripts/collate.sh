#!/bin/bash

# Directory with PNG files
input_dir="./output/$1"
output_pdf="$1.pdf"
temp_dir="./temp"
mkdir -p "$temp_dir"

# Dimensions for A4 in pixels at 150 DPI
a4_width=2480
a4_height=3508

# Effective content area with margins
margin=160
content_width=$((a4_width - 2 * margin))
content_height=$((a4_height - 2 * margin))

# Combine images two per A4 page
image_list=($(ls "$input_dir"/*.png))
for ((i=0; i<${#image_list[@]}; i+=2)); do
  img1="${image_list[i]}"
  img2="${image_list[i+1]}"
  
  # If the second image doesn't exist (odd number of images), use a placeholder
  if [[ -z "$img2" ]]; then
    magick -size ${content_width}x$((content_height/2)) xc:white "$temp_dir/placeholder.png"
    img2="$temp_dir/placeholder.png"
  fi

  # Create an A4 canvas with margins and ensure images fit within the content area
  magick -size ${a4_width}x${a4_height} xc:white \
    \( "$img1" -resize ${content_width}x$((content_height/2)) \) -gravity north -geometry +0+${margin} -composite \
    \( "$img2" -resize ${content_width}x$((content_height/2)) \) -gravity south -geometry +0+${margin} -composite \
    "$temp_dir/page_$((i/2 + 1)).png"
done

# Combine all pages into a single PDF
magick -density 300 "$temp_dir/page_*.png" -quality 100 "$output_pdf"

# Cleanup
rm -r "$temp_dir"

echo "PDF generated: $output_pdf"

