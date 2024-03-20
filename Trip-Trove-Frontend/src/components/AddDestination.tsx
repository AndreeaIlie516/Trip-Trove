import { useState } from "react";
import { Button, TextField } from "@mui/material";
import { IDestination } from "../interfaces/Destination";
import { useDestinations } from "../contexts/DestinationContext";
import { useNavigate } from "react-router-dom";

export function AddDestination() {
  const { addDestination, destinations } = useDestinations();
  const navigate = useNavigate();
  const [destination, setDestination] = useState<IDestination>({
    id: "",
    name: "",
    location_id: 0,
    image_url: "",
    description: "",
    is_private: false,
  });

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    destination.id = (destinations.length + 1).toString();
    console.log(destination);
    addDestination(destination);
    navigate("/destinations");
  };

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    setDestination((prevDestination) => ({
      ...prevDestination,
      [name]: name === "age" ? parseInt(value, 10) : value,
    }));
  };

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
        type="number"
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
        Add Destination
      </Button>
    </form>
  );
}