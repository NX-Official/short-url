export async function getShortenURL(url) {
    try {
        const res = await fetch('http://127.0.0.1:8080/api/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ url }),
        });

        if (!res.ok) {
            throw new Error(`HTTP error! status: ${res.status}`);
        }

        const result = await res.json();

        // 检查返回的代码是否表示成功
        if (result.code === 200) {
            // 成功，返回缩短后的 URL
            return result.data;
        } else {
            // 处理其他情况
            throw new Error(result.msg || 'Unknown error occurred');
        }
    } catch (error) {
        console.error('Error fetching shorten URL: ', error);
        throw error;
    }
}
