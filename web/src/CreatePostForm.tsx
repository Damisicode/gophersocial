import type React from "react";
import { useState } from "react";
import { API_URL } from "./App";

export const CreatePostForm: React.FC = () => {
    const [title, setTitle] = useState('')
    const [content, setContent] = useState('')

    const handleSubmit = async () => {
        await fetch(`${API_URL}/posts`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnb3BoZXJzb2NpYWwiLCJleHAiOjE3NjIyMDE5NjYsImlhdCI6MTc2MTk0Mjc2NiwiaXNzIjoiZ29waGVyc29jaWFsIiwibmJmIjoxNzYxOTQyNzY2LCJzdWIiOjIwN30.jiqi9gsYBlY08gCV2YRwnPMhRF_zMWoaMc04CweAZPA`
            },
            body: JSON.stringify({
                title,
                content
            })
        })

        setTitle('')
        setContent('')
    }

    return (
        <div className="gopher-form">
            <label>
                <input placeholder="Title..." value={title} type="text" onChange={(e) => setTitle(e.target.value)} />
            </label>
            <label>
                <textarea placeholder="What's in your mind..." value={content} onChange={(e) => setContent(e.target.value)} />
            </label>
            <button onClick={handleSubmit}>Share</button>
        </div>
    )
}