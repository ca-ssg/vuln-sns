import { defineStore } from 'pinia'
import axios from 'axios'

const API_URL = import.meta.env.VITE_API_URL

interface Post {
  id: number
  userId: string
  content: string
  createdAt: string
  likes: number
}

interface PostsState {
  posts: Post[]
  loading: boolean
  error: string | null
}

export const usePostsStore = defineStore('posts', {
  state: (): PostsState => ({
    posts: [],
    loading: false,
    error: null
  }),

  actions: {
    async fetchPosts() {
      this.loading = true
      try {
        const response = await axios.get(`${API_URL}/posts`)
        this.posts = response.data
      } catch (error) {
        console.error('Error fetching posts:', error)
        this.error = 'Failed to fetch posts'
      } finally {
        this.loading = false
      }
    },

    async createPost(content) {
      this.loading = true
      try {
        const response = await axios.post(`${API_URL}/posts`, { content })
        this.posts.unshift(response.data)
      } catch (error) {
        console.error('Error creating post:', error)
        this.error = 'Failed to create post'
      } finally {
        this.loading = false
      }
    },

    async updatePost(id, content) {
      this.loading = true
      try {
        await axios.put(`${API_URL}/posts/${id}`, { content })
        const index = this.posts.findIndex(post => post.id === id)
        if (index !== -1) {
          this.posts[index].content = content
        }
      } catch (error) {
        console.error('Error updating post:', error)
        this.error = 'Failed to update post'
      } finally {
        this.loading = false
      }
    },

    async deletePost(id) {
      this.loading = true
      try {
        await axios.delete(`${API_URL}/posts/${id}`)
        this.posts = this.posts.filter(post => post.id !== id)
      } catch (error) {
        console.error('Error deleting post:', error)
        this.error = 'Failed to delete post'
      } finally {
        this.loading = false
      }
    },

    async likePost(id) {
      try {
        await axios.post(`${API_URL}/posts/${id}/like`)
        const post = this.posts.find(p => p.id === id)
        if (post) {
          post.likes++
        }
      } catch (error) {
        console.error('Error liking post:', error)
        this.error = 'Failed to like post'
      }
    },

    async searchByHashtag(tag) {
      this.loading = true
      try {
        const response = await axios.get(`${API_URL}/search`, {
          params: { tag }
        })
        this.posts = response.data || []
      } catch (error) {
        console.error('Error searching posts:', error)
        this.error = 'Failed to search posts'
      } finally {
        this.loading = false
      }
    }
  }
})
