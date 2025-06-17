'use client'

import React, { useState, useEffect, useRef } from 'react'
import './chat-view.css'

function ChatView() {
  const [username, setUsername] = useState('')
  const [message, setMessage] = useState('')
  const [allMessages, setAllMessages] = useState<any[]>([])
  const [users, setUsers] = useState([])
  const [groups, setGroups] = useState<any[]>([])
  const [selectedUser, setSelectedUser] = useState('')
  const [groupName, setGroupName] = useState('')
  const [selectedGroupUsers, setSelectedGroupUsers] = useState([])
  const [showGroupForm, setShowGroupForm] = useState(false)
  const [isConnected, setIsConnected] = useState(false)
  const ws = useRef<WebSocket | null>(null)
  const messagesEndRef = useRef<any>(null)

  const filteredMessages = allMessages.filter((msg) => {
    if (!selectedUser) {
      return msg.to === ''
    }
    return (
      (msg.from === username && msg.to === selectedUser) ||
      (msg.to === username && msg.from === selectedUser) ||
      (msg.to === selectedUser && groups.includes(selectedUser))
    )
  })

  useEffect(() => {
    if (messagesEndRef.current) {
      messagesEndRef.current.scrollIntoView({ behavior: 'smooth' })
    }
  }, [filteredMessages])

  const connectToChat = () => {
    if (!username.trim()) return

    ws.current = new WebSocket('ws://localhost:8080/ws')

    ws.current.onopen = () => {
      ws.current!.send(
        JSON.stringify({
          type: 'connect',
          from: username,
        })
      )
      setIsConnected(true)
    }

    ws.current.onmessage = (event) => {
      const msg = JSON.parse(event.data)

      if (msg.type === 'user_list') {
        setUsers(msg.users.filter((u: string) => u !== username))
      } else if (msg.type === 'group_list' && Array.isArray(msg.users)) {
        setGroups((prev) => Array.from(new Set([...prev, ...msg.users])))
      } else if (msg.type === 'message') {
        setAllMessages((prev) => [...prev, msg])
      }
    }

    ws.current.onclose = () => {
      setIsConnected(false)
    }
  }

  const sendMessage = () => {
    if (!message.trim() || !username.trim()) return

    const msg = {
      type: 'message',
      from: username,
      to: selectedUser || '',
      content: message,
    }

    ws.current?.send(JSON.stringify(msg))
    setMessage('')
  }

  const createGroup = () => {
    if (!groupName.trim() || selectedGroupUsers.length === 0) return

    const msg = {
      type: 'create_group',
      from: username,
      to: groupName,
      users: selectedGroupUsers,
    }

    ws.current?.send(JSON.stringify(msg))
    setGroupName('')
    setSelectedGroupUsers([])
    setShowGroupForm(false)
  }

  const toggleUserSelection = (user: never) => {
    if (selectedGroupUsers.includes(user)) {
      setSelectedGroupUsers(selectedGroupUsers.filter((u) => u !== user))
    } else {
      setSelectedGroupUsers([...selectedGroupUsers, user])
    }
  }

  return (
    <div className="form-container">
      <h2>Go + React Chat</h2>

      {!isConnected ? (
        <div className="connect-form">
          <input
            type="text"
            placeholder="Enter your name"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <button onClick={connectToChat} disabled={!username}>
            Connect
          </button>
        </div>
      ) : (
        <div className="chat-container">
          <div className="user-list">
            <h3>Online Users</h3>
            <ul>
              <li className={!selectedUser ? 'selected' : ''} onClick={() => setSelectedUser('')}>
                Everyone
              </li>
              {users.map((user, idx) => (
                <li key={idx} className={selectedUser === user ? 'selected' : ''} onClick={() => setSelectedUser(user)}>
                  {user}
                </li>
              ))}
              {groups.map((group, idx) => (
                <li
                  key={idx}
                  className={selectedUser === group ? 'selected' : ''}
                  onClick={() => setSelectedUser(group)}
                >
                  {group} (group)
                </li>
              ))}
            </ul>
            <button onClick={() => setShowGroupForm(!showGroupForm)}>
              {showGroupForm ? 'Cancel' : 'Create Group'}
            </button>
            {showGroupForm && (
              <div className="group-form">
                <input
                  type="text"
                  placeholder="Group Name"
                  value={groupName}
                  onChange={(e) => setGroupName(e.target.value)}
                />
                <div className="group-users">
                  {users.map((user, idx) => (
                    <label key={idx}>
                      <input
                        type="checkbox"
                        checked={selectedGroupUsers.includes(user)}
                        onChange={() => toggleUserSelection(user)}
                      />
                      {user}
                    </label>
                  ))}
                </div>
                <button onClick={createGroup} disabled={!groupName || selectedGroupUsers.length === 0}>
                  Create
                </button>
              </div>
            )}
          </div>

          <div className="chat-box">
            <div className="messages">
              {filteredMessages.map((msg, idx) => (
                <div key={idx} className={`message ${msg.from === username ? 'sent' : 'received'}`}>
                  <div className="message-header">
                    <strong>{msg.from === username ? 'You' : msg.from}</strong>
                    <span className="message-time">{msg.time || ''}</span>
                  </div>
                  <div>{msg.content}</div>
                </div>
              ))}
              <div ref={messagesEndRef} />
            </div>

            <div className="message-input">
              <input
                type="text"
                placeholder={`Type a message${selectedUser ? ` to ${selectedUser}` : ''}...`}
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && sendMessage()}
              />
              <button onClick={sendMessage} disabled={!message}>
                Send
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default ChatView
