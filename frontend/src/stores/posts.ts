import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { axiosInstance as axios } from './auth'
import type { Post } from '@/types/post'

export const usePostsStore = defineStore('posts', () => {
  const posts = ref<Post[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const editModalVisible = ref(false)
  const currentPost = ref<Post | null>(null)

  const showEditModal = (post: Post) => {
    currentPost.value = { ...post }
    editModalVisible.value = true
  }

  const hideEditModal = () => {
    editModalVisible.value = false
    currentPost.value = null
  }

  const fetchPosts = async (): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      console.log('Fetching posts...')
      const response = await axios.get<Post[]>('/api/posts')
      console.log('Posts response:', response.data)
      posts.value = response.data
    } catch (err) {
      console.error('Error fetching posts:', err)
      error.value = 'Failed to fetch posts'
    } finally {
      loading.value = false
    }
  }

  const createPost = async (content: string): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      const response = await axios.post<Post>('/api/posts', { content })
      posts.value.unshift(response.data)
    } catch (err) {
      console.error('Error creating post:', err)
      error.value = 'Failed to create post'
    } finally {
      loading.value = false
    }
  }

  const updatePost = async (id: number, content: string): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      await axios.put(`/api/posts/${id}`, { content })
      const index = posts.value.findIndex((post: Post) => post.id === id)
      if (index !== -1) {
        posts.value[index].content = content
      }
    } catch (err) {
      console.error('Error updating post:', err)
      error.value = 'Failed to update post'
    } finally {
      loading.value = false
    }
  }

  const deletePost = async (id: number): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      await axios.delete(`/api/posts/${id}`)
      posts.value = posts.value.filter((post: Post) => post.id !== id)
    } catch (err) {
      console.error('Error deleting post:', err)
      error.value = 'Failed to delete post'
    } finally {
      loading.value = false
    }
  }

  const likePost = async (id: number): Promise<void> => {
    try {
      const response = await axios.post(`/api/posts/${id}/like`, {})
      const post = posts.value.find((p: Post) => p.id === id)
      if (post) {
        // Only increment likes if this is a new like (not a duplicate)
        if (response.status === 200 && !post.isLiked) {
          post.likes++
          post.isLiked = true
        }
      }
    } catch (err) {
      console.error('Error liking post:', err)
      error.value = 'Failed to like post'
    }
  }

  const searchByHashtag = async (tag: string): Promise<void> => {
    loading.value = true
    error.value = null
    try {
      const response = await axios.get<Post[]>('/api/search', {
        params: { tag }
      })
      
      // Ensure searchResults is always an array
      const searchResults = Array.isArray(response.data) ? response.data : []
      
      // 既存の投稿のisLikedプロパティを保持する
      const existingPosts = [...posts.value]
      posts.value = searchResults.map((newPost: Post) => {
        const existingPost = existingPosts.find((p: Post) => p.id === newPost.id)
        return existingPost ? { ...newPost, isLiked: existingPost.isLiked } : newPost
      })
    } catch (err) {
      console.error('Error searching posts:', err)
      error.value = 'Failed to search posts'
    } finally {
      loading.value = false
    }
  }

  const unlikePost = async (id: number): Promise<void> => {
    try {
      await axios.delete(`/api/posts/${id}/like`)
      const post = posts.value.find((p: Post) => p.id === id)
      if (post && post.isLiked) {
        post.likes--
        post.isLiked = false
      }
    } catch (err) {
      console.error('Error unliking post:', err)
      error.value = 'Failed to unlike post'
    }
  }

  return {
    posts,
    loading,
    error,
    editModalVisible,
    currentPost,
    fetchPosts,
    createPost,
    updatePost,
    deletePost,
    likePost,
    unlikePost,
    showEditModal,
    hideEditModal,
    searchByHashtag
  }
})
