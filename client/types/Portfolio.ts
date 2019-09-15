export interface SubCategory{
    id: string
    mainCategoryId: string
    name: string
    description: string
    createdAt: string
    updatedAt: string
}

export interface MainCategory{
    id: string
    name: string
    description: string
    subCategories: SubCategory[]
    createdAt: string
    updatedAt: string
}

export interface Content{
    id: string
    categoryId: string
    title: string
    image: string
    date: string
    createdAt: string
    updatedAt: string
}

export interface ContentDetail{
    id: string
    categoryId: string
    title: string
    image: string
    description: string
    date: string
    subImagesCount: number
    subImages: SubImage[]
    taggedContents: TaggedContent[]
    subCategory: SubCategory
    mainCategory: MainCategory
    createdAt: string
    updatedAt: string
}

export interface SubImage{
    id: string
    name: string
    contentId: string
    url: string
    createdAt: string
    updatedAt: string
}

export interface TaggedContent{
    id: string
    tagId: string
    contentId: string
    tag: Tag
    createdAt: string
    updatedAt: string
}

export interface Tag{
    id: string
    name: string
    description: string
    createdAt: string
    updatedAt: string
}