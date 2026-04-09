import React, { useState, useEffect, useRef } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import api from '../services/api';
import { Send, ArrowLeft, Ticket, MoreVertical, Loader2, Info, ChevronDown } from 'lucide-react';

const Chat = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [reply, setReply] = useState('');
  const [sending, setSending] = useState(false);
  const [showEscalate, setShowEscalate] = useState(false);
  const [ticketData, setTicketData] = useState({ title: '', desc: '', priority: 'medium' });
  const messagesEndRef = useRef(null);

  useEffect(() => {
    fetchDetail();
  }, [id]);

  useEffect(() => {
    scrollToBottom();
  }, [data]);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };

  const fetchDetail = async () => {
    try {
      const response = await api.get(`/conversations/${id}`);
      setData(response.data.data);
    } catch (err) {
      console.error(err);
      navigate('/');
    } finally {
      setLoading(false);
    }
  };

  const handleSendReply = async (e) => {
    e.preventDefault();
    if (!reply.trim() || sending) return;

    setSending(true);
    try {
      await api.post(`/conversations/${id}/messages`, { message: reply });
      setReply('');
      fetchDetail();
    } catch (err) {
      alert('Failed to send message');
    } finally {
      setSending(false);
    }
  };

  const handleEscalate = async (e) => {
    e.preventDefault();
    try {
      await api.post('/tickets', {
        conv_id: parseInt(id),
        ...ticketData
      });
      alert('Conversation escalated to ticket!');
      setShowEscalate(false);
    } catch (err) {
      alert(err.response?.data?.message || 'Escalation failed');
    }
  };

  if (loading) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <Loader2 className="animate-spin text-primary" size={40} />
      </div>
    );
  }

  const { conversation, messages } = data;

  return (
    <div className="flex-1 flex overflow-hidden space-y-8">
      {/* Chat Area */}
      <div className="flex-1 flex flex-col bg-bg-dark" style={{ borderRight: '1px solid var(--border)' }}>
        {/* Header */}
        <div className="p-4 flex items-center justify-between border-b" style={{ borderColor: 'var(--border)' }}>
          <div className="flex items-center gap-4">
            <button className="ghost p-2 rounded-full" onClick={() => navigate('/')}>
              <ArrowLeft size={20} />
            </button>
            <div>
              <h3 style={{ fontWeight: '600' }}>Customer {conversation.customer_external_id}</h3>
              <span className="text-xs text-success" style={{ color: 'var(--success)' }}>Active Chat</span>
            </div>
          </div>
          <button
            className="flex items-center gap-2 px-3 py-1.5 text-xs text-white mb-32"
            onClick={() => setShowEscalate(!showEscalate)}
          >
            <Ticket size={14} /> {showEscalate ? 'Close Escalation' : 'Escalate to Ticket'}
          </button>
        </div>

        {/* Messages */}
        <div className="flex-1 overflow-y-auto p-6 space-y-6 mt-8">
          {messages.map((msg) => (
            <div
              key={msg.id}
              className={`flex flex-col ${msg.sender_type === 'agent' ? 'items-end' : 'items-start'}`}
              style={{ marginBottom: '24px' }}
            >
              <div
                className={`p-4 max-w-[70%] rounded-2xl ${msg.sender_type === 'agent'
                  ? 'bg-primary text-white rounded-tr-none'
                  : 'glass rounded-tl-none'
                  }`}
                style={{
                  backgroundColor: msg.sender_type === 'agent' ? 'var(--primary)' : 'rgba(255,255,255,0.05)',
                  animation: 'fadeIn 0.2s ease-out'
                }}
              >
                <p style={{ fontSize: '15px' }}>{msg.message}</p>
                <span className="block mt-1 text-[10px] opacity-60">
                  {new Date(msg.created_at).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                </span>
              </div>
            </div>
          ))}
          <div ref={messagesEndRef} />
        </div>

        {/* Input */}
        <div className="p-4 border-t" style={{ borderColor: 'var(--border)' }}>
          <form onSubmit={handleSendReply} className="flex gap-3">
            <input
              placeholder="Type your reply..."
              value={reply}
              onChange={(e) => setReply(e.target.value)}
              className="flex-1"
            />
            <button
              type="submit"
              className="primary p-3 flex items-center justify-center rounded-xl"
              disabled={sending || !reply.trim()}
            >
              {sending ? <Loader2 className="animate-spin" size={20} /> : <Send size={20} />}
            </button>
          </form>
        </div>
      </div>

      {/* Escalation Sidebar */}
      {showEscalate && (
        <div className="w-[350px] bg-bg-card p-6 overflow-y-auto fade-in">
          <div className="flex items-center gap-2 mb-6 text-primary">
            <Ticket size={24} />
            <h2 style={{ fontSize: '20px', fontWeight: '700' }}>Escalate Ticket</h2>
          </div>

          <form onSubmit={handleEscalate} className="space-y-6">
            <div style={{ marginBottom: '16px' }}>
              <label className="block text-xs font-semibold mb-2 uppercase opacity-50">Ticket Title</label>
              <input
                placeholder="Issue with login..."
                required
                value={ticketData.title}
                onChange={(e) => setTicketData({ ...ticketData, title: e.target.value })}
              />
            </div>

            <div style={{ marginBottom: '16px' }}>
              <label className="block text-xs font-semibold mb-2 uppercase opacity-50">Priority</label>
              <select
                className="w-full bg-bg-input border border-border p-3 rounded-lg text-white"
                style={{ backgroundColor: 'var(--bg-input)', border: '1px solid var(--border)', color: 'white', padding: '12px' }}
                value={ticketData.priority}
                onChange={(e) => setTicketData({ ...ticketData, priority: e.target.value })}
              >
                <option value="low">Low</option>
                <option value="medium">Medium</option>
                <option value="high">High</option>
                <option value="urgent">Urgent</option>
              </select>
            </div>

            <div style={{ marginBottom: '24px' }}>
              <label className="block text-xs font-semibold mb-2 uppercase opacity-50">Description</label>
              <textarea
                className="w-full bg-bg-input border border-border p-3 rounded-lg text-white min-h-[120px]"
                style={{ backgroundColor: 'var(--bg-input)', border: '1px solid var(--border)', color: 'white', padding: '12px', minHeight: '120px' }}
                placeholder="Details about the escalation..."
                required
                value={ticketData.desc}
                onChange={(e) => setTicketData({ ...ticketData, desc: e.target.value })}
              />
            </div>

            <button className="primary w-full" type="submit">
              Raise Ticket
            </button>
          </form>

          <div className="mt-8 p-4 rounded-xl bg-primary/5 border border-primary/20 text-xs">
            <div className="flex items-center gap-2 mb-2 text-primary">
              <Info size={14} />
              <span className="font-semibold">Note</span>
            </div>
            <p className="opacity-70 leading-relaxed">
              Escalating this conversation will create a formal ticket in the support system indexed under this tenant.
            </p>
          </div>
        </div>
      )}
    </div>
  );
};

export default Chat;
