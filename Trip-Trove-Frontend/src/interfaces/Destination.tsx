export interface Destination {
    id: number,
    name: string,
    location?: string,
    country?: string,
    image_url: string,
    description?: string,
    is_private: boolean
}

export interface DestinationCardProps {
    data: Destination;
}