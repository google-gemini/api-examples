set -eu

echo "[START grounding_maps]"
# [START grounding_maps]"
curl -X POST 'https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash-lite:generateContent' \
  -H 'Content-Type: application/json' \
  -H "x-goog-api-key: ${GEMINI_API_KEY}" \
  -d '{
  "contents": [{
    "role": "user",
    "parts": [{
      "text": "What are the best Italian restaurants within a 15-minute walk from here?"
    }]
  }],
  "tools": [{"googleMaps": {}}],
  "toolConfig": {
    "retrievalConfig": {
      "latLng": {"latitude": 34.050481, "longitude": -118.248526}
    }
  }
}'

# [END grounding_maps]
