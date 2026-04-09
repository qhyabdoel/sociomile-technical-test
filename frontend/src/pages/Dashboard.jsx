import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';
import { MessageCircle, Clock, ChevronRight, Loader2, User } from 'lucide-react';

const Dashboard = () => {
  const [conversations, setConversations] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    fetchConversations();
  }, []);

  const fetchConversations = async () => {
    try {
      const response = await api.get('/conversations');
      setConversations(response.data.data || []);
    } catch (err) {
      setError('Failed to load conversations.');
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const getStatusColor = (status) => {
    return status === 'open' ? 'var(--success)' : 'var(--text-secondary)';
  };

  if (loading) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <Loader2 className="animate-spin text-primary" size={40} />
      </div>
    );
  }

  return (
    <div className="p-8 fade-in">
      <div className="flex justify-between items-center mb-8">
        <div>
          <h1 style={{ fontSize: '28px', fontWeight: '700', marginBottom: '8px' }}>Conversations</h1>
          <p style={{ color: 'var(--text-secondary)' }}>Manage your active tenant discussions</p>
        </div>
        <div className="glass px-4 py-2 flex items-center gap-2" style={{ borderRadius: '20px' }}>
          <div className="w-2 h-2 rounded-full" style={{ backgroundColor: 'var(--success)' }}></div>
          <span style={{ fontSize: '14px', fontWeight: '500' }}>{conversations.length} Active</span>
        </div>
      </div>

      {error && (
        <div className="bg-error/10 text-error p-4 rounded-lg mb-6">
          {error}
        </div>
      )}

      {conversations.length === 0 ? (
        <div className="glass p-12 text-center">
          <MessageCircle size={48} className="mx-auto mb-4 opacity-20" style={{ margin: '0 auto 16px' }} />
          <h3 style={{ fontSize: '18px', fontWeight: '600' }}>No conversations yet</h3>
          <p style={{ color: 'var(--text-secondary)' }}>New interactions will appear here.</p>
        </div>
      ) : (
        <div className="grid gap-4">
          {conversations.map((conv) => (
            <div 
              key={conv.id}
              onClick={() => navigate(`/conversations/${conv.id}`)}
              className="glass p-5 flex items-center justify-between cursor-pointer hover:border-primary/50 transition-all hover:translate-x-1"
              style={{ transition: 'all 0.2s', border: '1px solid var(--glass-border)' }}
            >
              <div className="flex items-center gap-4">
                <div className="w-12 h-12 rounded-full bg-primary/10 flex items-center justify-center" style={{ backgroundColor: 'rgba(99, 102, 241, 0.1)' }}>
                  <User size={24} className="text-primary" style={{ color: 'var(--primary)' }} />
                </div>
                <div>
                  <h3 style={{ fontWeight: '600', marginBottom: '4px' }}>Customer {conv.customer_external_id}</h3>
                  <div className="flex items-center gap-3 text-xs" style={{ color: 'var(--text-secondary)' }}>
                    <span className="flex items-center gap-1">
                      <Clock size={14} /> {new Date(conv.created_at).toLocaleDateString()}
                    </span>
                    <span className="flex items-center gap-1">
                      <div className="w-1.5 h-1.5 rounded-full" style={{ backgroundColor: getStatusColor(conv.status) }}></div>
                      {conv.status.toUpperCase()}
                    </span>
                  </div>
                </div>
              </div>
              <ChevronRight size={20} style={{ color: 'var(--text-secondary)' }} />
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Dashboard;
