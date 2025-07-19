// src/api/aichatCen.js
import OpenAI from 'openai';

const modelConfig = {
  'deepseek-chat': {
    baseURL: 'https://api.deepseek.com/beta',
    apiKey: import.meta.env.VITE_DEEPSEEK_API_KEY,
    model: 'deepseek-chat',
  }
};

const openai = new OpenAI({
  baseURL: modelConfig['deepseek-chat'].baseURL,
  apiKey: modelConfig['deepseek-chat'].apiKey,
  dangerouslyAllowBrowser: true // ⚠️ 仅限开发测试，生产环境必须用后端代理
});

export async function callDeepSeekAPI(prompt, chatHistory) {
  try {
    const response = await openai.chat.completions.create({
      model: 'deepseek-chat',
      messages: [
        {
          role: "system",
          content: "你是专业的医学影像AI助手"
        },
        ...chatHistory.map(msg => ({
          role: msg.from === 'user' ? 'user' : 'assistant',
          content: msg.text
        }))
      ],
      temperature: 0.7,
      max_tokens: 2000
    });
    console.log('回复：', response.choices[0]?.message?.content);
    return {
      success: true,
      content: response.choices[0]?.message?.content || '未能获取有效回复'
    };
  } catch (error) {
    console.error('API调用错误:', error);
    return {
      success: false,
      content: '抱歉，获取回答时出错，请稍后再试。'
    };
  }
}