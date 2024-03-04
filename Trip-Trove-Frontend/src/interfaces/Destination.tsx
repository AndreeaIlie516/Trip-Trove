export interface Destination {
    id: number,
    name: string,
    locationId: number,
    image_url: string,
    description?: string,
    is_private: boolean
}

export interface DestinationCardProps {
    data: Destination;
}