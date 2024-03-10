export interface IDestination {
    id: number,
    name: string,
    location_id: number,
    image_url: string,
    description?: string,
    is_private: boolean
}

export interface IDestinationCardProps {
    data: IDestination;
}