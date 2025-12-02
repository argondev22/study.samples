from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/api/health', methods=['GET'])
def health_check():
    return jsonify({
        'status': 'healthy',
        'message': 'APIサーバーは正常に動作しています'
    })

@app.route('/api/hello', methods=['GET'])
def hello():
    return jsonify({
        'message': 'こんにちは！'
    })

if __name__ == '__main__':
    app.run(debug=True) 