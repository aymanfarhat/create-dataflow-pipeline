python main.py \
  --runner=DirectRunner \
  --staging_location=$STAGING_LOCATION \
  --temp_location=$TEMP_LOCATION \
  --setup_file=./setup.py \
  --input_file=$INPUT_FILE \
  --output_bq_table=$OUTPUT_BQ_TABLE
