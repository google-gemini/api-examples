set -eu

echo "[START embed_content]"
# [START embed_content]
curl "https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-001:embedContent" \
-H "x-goog-api-key: $GEMINI_API_KEY" \
-H 'Content-Type: application/json' \
-d '{"model": "models/gemini-embedding-001",
     "content": {
     "parts":[{
     "text": "What is the meaning of life?"}]}
    }'
# [END embed_content]

echo "[START batch_embed_contents]"
# [START batch_embed_contents]
curl "https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-001:batchEmbedContents" \
-H "x-goog-api-key: $GEMINI_API_KEY" \
-H 'Content-Type: application/json' \
-d '{"requests": [{
      "model": "models/gemini-embedding-001",
      "content": {
      "parts":[{
        "text": "What is the meaning of life?"}]}, },
      {
      "model": "models/gemini-embedding-001",
      "content": {
      "parts":[{
        "text": "How much wood would a woodchuck chuck?"}]}, },
      {
      "model": "models/gemini-embedding-001",
      "content": {
      "parts":[{
        "text": "How does the brain work?"}]}, }, ]}' 2> /dev/null | grep -C 5 values
# [END batch_embed_contents]