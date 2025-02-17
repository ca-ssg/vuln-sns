import { defineStore } from 'pinia'
import { ref } from 'vue'
import axios from 'axios'
import { useAuthStore } from './auth'

interface Post {
  id: number
  userId: string
  content: string
  createdAt: string
  updatedAt: string
  likes: number
  isLiked?: boolean
}

export const usePostStore = defineStore('posts', () => {
  const posts = ref<Post[]>([])
  const authStore = useAuthStore()

  const fetchPosts = async () => {
    try {
      const response = await axios.get(`${import.meta.env.VITE_API_URL}/posts`)
      posts.value = response.data
    } catch (error) {
      console.error('Failed to fetch posts:', error)
    }
  }

  const createPost = async (content: string) => {
    if (!authStore.token) return false

    try {
      const response = await axios.post(
        `${import.meta.env.VITE_API_URL}/posts`,
        { content },
        {
          headers: { Authorization: `Bearer ${authStore.token}` }
        }
      )
      await fetchPosts()
      return true
    } catch (error) {
      console.error('Failed to create post:', error)
      return false
    }
  }

  const updatePost = async (id: number, content: string) => {
    if (!authStore.token) return false

    try {
      await axios.put(
        `http://localhost:8080/posts/${id}`,
        { content },
        {
          headers: { Authorization: `Bearer ${authStore.token}` }
        }
      )
      await fetchPosts()
      return true
    } catch (error) {
      console.error('Failed to update post:', error)
      return false
    }
  }

  const deletePost = async (id: number) => {
    if (!authStore.token) return false

    try {
      await axios.delete(`http://localhost:8080/posts/${id}`, {
        headers: { Authorization: `Bearer ${authStore.token}` }
      })
      await fetchPosts()
      return true
    } catch (error) {
      console.error('Failed to delete post:', error)
      return false
    }
  }

  const toggleLike = async (id: number) => {
    if (!authStore.token) return false

    try {
      await axios.post(
        `http://localhost:8080/posts/${id}/like`,
        {},
        {
          headers: { Authorization: `Bearer ${authStore.token}` }
        }
      )
      await fetchPosts()
      return true
    } catch (error) {
      console.error('Failed to toggle like:', error)
      return false
    }
  }

  return {
    posts,
    fetchPosts,
    createPost,
    updatePost,
    deletePost,
    toggleLike
  }
})
