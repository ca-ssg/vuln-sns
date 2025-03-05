export interface Post {
  id: number
  userId: string
  content: string
  createdAt: string
  likes: number
  isLiked: boolean
  user_id?: string // For backward compatibility
}
