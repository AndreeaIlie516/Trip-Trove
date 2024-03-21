export interface IDestination {
    id: string,
    name: string,
    location_id: number,
    image_url: string,
    visitors_last_year: number,
    description?: string,
    is_private: boolean
}

export interface IDestinationCardProps {
    data: IDestination;
}