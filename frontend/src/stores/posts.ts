import { defineStore } from 'pinia'
import { axiosInstance as axios } from './auth'

interface Post {
  id: number
  userId: string
  content: string
  createdAt: string
  likes: number
  isLiked: boolean
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
    async fetchPosts(): Promise<void> {
      this.loading = true
      try {
        const response = await axios.get<Post[]>('/posts')
        this.posts = response.data
      } catch (error) {
        console.error('Error fetching posts:', error)
        this.error = 'Failed to fetch posts'
      } finally {
        this.loading = false
      }
    },

    async createPost(content: string): Promise<void> {
      this.loading = true
      try {
        const response = await axios.post<Post>('/posts', { content })
        this.posts.unshift(response.data)
      } catch (error) {
        console.error('Error creating post:', error)
        this.error = 'Failed to create post'
      } finally {
        this.loading = false
      }
    },

    async updatePost(id: number, content: string): Promise<void> {
      this.loading = true
      try {
        await axios.put(`/posts/${id}`, { content })
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

    async deletePost(id: number): Promise<void> {
      this.loading = true
      try {
        await axios.delete(`/posts/${id}`)
        this.posts = this.posts.filter(post => post.id !== id)
      } catch (error) {
        console.error('Error deleting post:', error)
        this.error = 'Failed to delete post'
      } finally {
        this.loading = false
      }
    },

    async likePost(id: number): Promise<void> {
      try {
        const response = await axios.post(`/posts/${id}/like`, {})
        const post = this.posts.find(p => p.id === id)
        if (post) {
          // Only increment likes if this is a new like (not a duplicate)
          if (response.status === 200 && !post.isLiked) {
            post.likes++
          }
          post.isLiked = true
        }
      } catch (error) {
        console.error('Error liking post:', error)
        this.error = 'Failed to like post'
      }
    },

    async searchByHashtag(tag: string): Promise<void> {
      this.loading = true
      try {
        const response = await axios.get<Post[]>('/search', {
          params: { tag }
        })
        this.posts = response.data || []
      } catch (error) {
        console.error('Error searching posts:', error)
        this.error = 'Failed to search posts'
      } finally {
        this.loading = false
      }
    },

    async unlikePost(id: number): Promise<void> {
      try {
        await axios.delete(`/posts/${id}/like`)
        const post = this.posts.find(p => p.id === id)
        if (post && post.isLiked) {
          post.likes--
          post.isLiked = false
        }
      } catch (error) {
        console.error('Error unliking post:', error)
        this.error = 'Failed to unlike post'
      }
    }
  }
})
