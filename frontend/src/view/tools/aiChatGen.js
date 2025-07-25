// src/api/aichatCen.js
import OpenAI from 'openai';
import myRAG from './myRAG.json';// 确保路径正确
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
  dangerouslyAllowBrowser: true //仅限开发测试，生产环境必须用后端代理
});
const systemPrompts = myRAG.map(item => `${item.标题}: ${item.概述}`).join('\n\n');
export async function callDeepSeekAPI(prompt, chatHistory) {
  try {
    const response = await openai.chat.completions.create({
      model: 'deepseek-chat',
      messages: [
        {
          role: "system",
          content: `你是集成在一个超声影像分级网站的ai助手——分级喵，以下是记忆仓库中存储的本网站中保存的针对各种医学影像的分级方案：\n\n${systemPrompts}，请你不要回答用户询问的和超声影像分级相关专业知识之外的任何问题！在回答用户问题的过程中，向他推荐最适合他的存储在本记忆仓库中的解决方案。`
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