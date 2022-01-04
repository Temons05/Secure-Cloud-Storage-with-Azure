type JWTToken = string

interface LoginResponse {
    success: boolean
    data: {
        token: JWTToken
    }
}

interface RegisterResponse {
    success?: boolean
    error?: string

}

interface UploadResponse {
    success: boolean
}

interface AzureFile {
    file_name: string
    last_modified: number
    md5: string
    size: number
    uid: string
    version: string
}

interface AzureFileParsed {
    [key: string]: AzureFile[]

}
interface BlobListResponse {
    success: boolean
    data: AzureFile[]
}

interface HTMLInputFileEvent {
    target: {
        files: FileList
    }
}