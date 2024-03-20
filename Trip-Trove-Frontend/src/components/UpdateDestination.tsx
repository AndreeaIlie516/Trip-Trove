import { useEffect, useState } from "react";
import { Button, TextField } from "@mui/material";
import { IDestination } from "../interfaces/Destination";
import { useDestinations } from "../contexts/DestinationContext";
import { useNavigate, useParams } from "react-router-dom";

export function UpdateDestination() {
  const { getDestinationById, updateDestination } = useDestinations();
  const navigate = useNavigate();
  const { id } = useParams();
  const [destination, setDestination] = useState<IDestination | undefined>(
    undefined
  );

  useEffect(() => {
    const fetchDestinationDetails = async () => {
      if (!id) {
        return null;
      }
      const fetchedDestination = await getDestinationById(parseInt(id));
      setDestination(fetchedDestination);
    };
    fetchDestinationDetails();
  }, []);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!destination) {
      return;
    }
    updateDestination(destination);
    navigate("/destinations");
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setDestination((prevDestination) =>
      prevDestination ? { ...prevDestination, [name]: value } : prevDestination
    );
  };

  if (!destination) return <div>Loading...</div>;

  return (
    <form onSubmit={handleSubmit}>
      <TextField
        margin="normal"
        fullWidth
        label="Name"
        name="name"
        value={destination.name}
        onChange={handleChange}
      />
      <TextField
        margin="normal"
        fullWidth
        label="Location ID"
        name="location_id"
        value={destination.location_id}
        onChange={handleChange}
      />
      <TextField
        margin="normal"
        fullWidth
        label="Image Url"
        name="image_url"
        value={destination.image_url}
        onChange={handleChange}
      />
      <TextField
        margin="normal"
        fullWidth
        label="Description"
        name="description"
        value={destination.description}
        onChange={handleChange}
      />

      <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
        Update Destination
      </Button>
    </form>
  );
}
