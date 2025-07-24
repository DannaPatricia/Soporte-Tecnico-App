export default function (apiClient) {
  return {
    async getMessagesByCase(supportCaseId) {
      return apiClient.post('/messages-by-case', { supportCaseId })
    },
    async sendMessage(newMessage) {
      return apiClient.post('/create-messages', {
        content: newMessage.content,
        sent_at: newMessage.sent_at,
        supportCaseId: newMessage.supportCaseId,
        userId: newMessage.userId,
        sender: newMessage.sender
      })
    },
  }
}
