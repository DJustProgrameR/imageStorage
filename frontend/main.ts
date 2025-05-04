import express from 'express';
import {PetApi, GetPet200Response, Configuration} from './src/client';
import axios from 'axios';
import path from "path";

const app = express();
const port = process.env.PORT || 8000;

// Configure API client
const apiConfig = new Configuration({
    basePath: process.env.API_BASE_URL || 'http://localhost:8080',
    baseOptions: {
        timeout: 10000,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        }
    }
});


const petApi = new PetApi(apiConfig);

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));

app.use(express.static(path.join(__dirname, 'public')));

// Routes
app.get('/', async (req, res) => {
    try {
        // Fetch current pet
        const response = await petApi.getPet();
        const pet = response.data;

        const rows = countRows(pet);
        const cols = countCols(pet);

        res.render('index', {
            rows: rows,
            cols: cols,
            title: 'Сервис питомцев',
            pet: pet
        });
    } catch (error) {
        console.error('Failed to fetch pet:', error);
        res.status(500).render('error', {
            message: 'Failed to load pet data'
        });
    }
});

app.post('/update-pet', async (req, res) => {
    try {
        const { ascii, description } = req.body;

        if (!ascii || !description) {
            return res.status(400).json({ error: 'Both ASCII art and description are required' });
        }

        // Update pet
        await petApi.uploadPet({ ascii, description });
        res.redirect('/');
    } catch (error) {
        console.error('Failed to update pet:', error);
        res.status(500).json({ error: 'Failed to update pet' });
    }
});

// Error handling middleware
app.use((err: Error, req: express.Request, res: express.Response, next: express.NextFunction) => {
    console.error('Unhandled error:', err);
    res.status(500).render('error', {
        message: 'Something went wrong'
    });
});

// Start server
app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
    console.log(`API base URL: ${apiConfig.basePath}`);
});

// Graceful shutdown
process.on('SIGTERM', () => {
    console.log('SIGTERM received. Shutting down gracefully');
    process.exit(0);
});

process.on('SIGINT', () => {
    console.log('SIGINT received. Shutting down gracefully');
    process.exit(0);
});

function countRows(pet?:  GetPet200Response): number {
    if(!pet) {
        return 10
    }
    const newlineRegex = /\n/g;
    const matches = pet.ascii.match(newlineRegex);
    return matches ? Math.max(10,matches.length+4) : 0;
}

function countCols(pet?:  GetPet200Response): number {
    if(!pet) {
        return 50
    }
    let s = pet.ascii;
    let strings = s.split('\n');
    let n = strings.length;
    let maxLength = 0;

    for (let i = 0; i < n; i++) {
        maxLength = Math.max(maxLength, strings[i].length);
    }

    return Math.max(50,maxLength*1.5);
}