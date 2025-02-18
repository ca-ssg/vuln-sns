import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Post {
  id: number
  userId: string
  content: string
  createdAt: string
  updatedAt: string
  likes: number
}

export const usePostStore = defineStore('posts', () => {
  const posts = ref<Post[]>([])

  const fetchPosts = async () => {
    try {
      const response = await fetch('http://localhost:9090/api/posts')
      if (!response.ok) {
        throw new Error('Failed to fetch posts')
      }
      posts.value = await response.json()
    } catch (error) {
      console.error('Error fetching posts:', error)
    }
  }

  const createPost = async (content: string) => {
    try {
      const token = localStorage.getItem('token')
      const user = JSON.parse(localStorage.getItem('user') || '{}')
      
      const response = await fetch('http://localhost:9090/api/posts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': token || '',
        },
        body: JSON.stringify({
          userId: user.id,
          content,
        }),
      })

      if (!response.ok) {
        throw new Error('Failed to create post')
      }

      const newPost = await response.json()
      posts.value.unshift(newPost)
    } catch (error) {
      console.error('Error creating post:', error)
    }
  }

  const toggleLike = async (postId: number) => {
    try {
      const token = localStorage.getItem('token')
      const response = await fetch(`http://localhost:9090/api/posts/${postId}/like`, {
        method: 'POST',
        headers: {
          'Authorization': token || '',
        },
      })

      if (!response.ok) {
        throw new Error('Failed to toggle like')
      }

      const post = posts.value.find(p => p.id === postId)
      if (post) {
        post.likes = post.likes + 1
      }
    } catch (error) {
      console.error('Error toggling like:', error)
    }
  }

  return {
    posts,
    fetchPosts,
    createPost,
    toggleLike,
  }
})
