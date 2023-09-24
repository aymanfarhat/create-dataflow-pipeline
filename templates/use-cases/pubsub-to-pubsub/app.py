import re
from typing import Optional
import apache_beam as beam
from apache_beam.options.pipeline_options import PipelineOptions



def run(
    input_file: str,
    output_bq_table: str,
    beam_options: Optional[PipelineOptions] = None
) -> None:
    with beam.Pipeline(options=beam_options) as pipeline:
        schema = {
            'fields': [
                {'name': 'word', 'type': 'STRING', 'mode': 'NULLABLE'},
                {'name': 'count', 'type': 'INTEGER', 'mode': 'NULLABLE'},
            ]
        }

        word_count = (
            pipeline
             | 'Read file' >> beam.io.ReadFromText(input_file)
             | 'Flatten words' >> beam.FlatMap(lambda line: re.findall(r'[\w\']+', line.strip(), re.UNICODE))
             | 'Count words' >> beam.combiners.Count.PerElement()
             | 'Format results' >> beam.MapTuple(lambda word, count: {'word': word, 'count': count})
        )

        word_count | 'Write to BigQuery' >> beam.io.WriteToBigQuery(
            table=output_bq_table,
            schema=schema,
            create_disposition=beam.io.BigQueryDisposition.CREATE_IF_NEEDED,
            write_disposition=beam.io.BigQueryDisposition.WRITE_TRUNCATE,
        )