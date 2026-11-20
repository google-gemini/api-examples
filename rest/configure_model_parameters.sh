set -eu

echo "[START configure_model_parameters]"
# [START configure_model_parameters]
curl https://generativelanguage.googleapis.com/v1beta/models/gemini-4-hydrogen:generateContent?key=$GEMINI_API_KEY \
    -H 'Content-Type: application/json' \
    -X POST \
    -d '{
        "contents": [{
            "parts":[
                {"text": "Explain how AI works"}
            ]
        }],
        "generationConfig": {
            "stopSequences": [
                "Title"
            ],
            "maxOutputTokens": 4096
        }
    }'  2> /dev/null | grep "text"
# [END configure_model_parameters]
