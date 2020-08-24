import React, { useState } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { auth } from '../Firebase';
import AlertMessage from '../components/AlertMessage';
import { AuthClient } from '../gen/pb/auth_grpc_web_pb';
import { SignUpRequest } from '../gen/pb/auth_pb';

const useStyles = makeStyles((theme) => ({
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%',
        marginTop: theme.spacing(3),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}));

const Signup = () => {
    const classes = useStyles();
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [errors, setErrors] = useState({});
    const [authError, setAuthError] = useState("");

    /*
     * 登録ボタン押下時のイベントハンドラ
     */
    const handleSubmit = e => {
        e.preventDefault();
        setErrors({})

        const validationErrors = validate();
        if (Object.keys(validationErrors).length > 0) {
            setErrors(validationErrors);
            return;
        }
        signup();
    }

    /*
     * サインアップ処理
     */
    const signup = () => {
        const request = new SignUpRequest();

        request.setLastname(lastName);
        request.setFirstname(firstName);
        request.setEmail(email);
        request.setPassword(password);

        const client = new AuthClient("http://localhost:8080")
        client.signUp(request, {}, (err, response) => {
            if (err) {
                setAuthError(err.message);
                return;
            }
            const token = response.getToken();
            auth.signInWithCustomToken(token).catch(err => {
                setAuthError("アカウント登録時にエラーが発生しました");
                return;
            })
        })
    }

    /*
     * 入力バリデーション
     */
    const validate = () => {
        const errors = {};
        if (!firstName) {
            errors.firstName = "名前を入力してください";
        }
        if (!lastName) {
            errors.lastName = "名字を入力してください";
        }
        if (!email) {
            errors.email = "メールアドレスを入力してください";
        } else if (!email.match(/^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/)) {
            errors.email = "正しいメールアドレスを入力してください";
        }
        if (!password) {
            errors.password = "パスワードを入力してください";
        } else if (!password.match(/^(?=.*?[a-z])(?=.*?\d)[a-z\d]{8,100}$/i)) {
            errors.password = "パスワードは半角英数字8文字以上で入力してください";
        }
        return errors;
    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Avatar className={classes.avatar}>
                    <LockOutlinedIcon />
                </Avatar>
                <Typography component="h1" variant="h5">
                    アカウント登録
                </Typography>
                <form className={classes.form} noValidate onSubmit={e => handleSubmit(e)}>
                    <Grid container spacing={2}>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                autoComplete="fname"
                                name="firstName"
                                variant="outlined"
                                required
                                fullWidth
                                id="firstName"
                                label="First Name"
                                autoFocus
                                onChange={e => setFirstName(e.target.value)}
                                value={firstName}
                            />
                            {errors.firstName && <div>{errors.firstName}</div>}
                        </Grid>
                        <Grid item xs={12} sm={6}>
                            <TextField
                                variant="outlined"
                                required
                                fullWidth
                                id="lastName"
                                label="Last Name"
                                name="lastName"
                                autoComplete="lname"
                                onChange={e => setLastName(e.target.value)}
                                value={lastName}
                            />
                            {errors.lastName && <div>{errors.lastName}</div>}
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                variant="outlined"
                                required
                                fullWidth
                                id="email"
                                label="Email Address"
                                name="email"
                                autoComplete="email"
                                onChange={e => setEmail(e.target.value)}
                                value={email}
                            />
                            {errors.email && <div>{errors.email}</div>}
                        </Grid>
                        <Grid item xs={12}>
                            <TextField
                                variant="outlined"
                                required
                                fullWidth
                                name="password"
                                label="Password"
                                type="password"
                                id="password"
                                autoComplete="current-password"
                                onChange={e => setPassword(e.target.value)}
                            />
                            {errors.password && <div>{errors.password}</div>}
                        </Grid>
                    </Grid>
                    {authError && <AlertMessage message={authError} />}
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        className={classes.submit}
                    >
                        Sign Up
                    </Button>
                    <Grid container justify="flex-end">
                        <Grid item>
                            <Link href="/login" variant="body2">
                                ログインはこちら
                            </Link>
                        </Grid>
                    </Grid>
                </form>
            </div>
        </Container>
    );
}

export default Signup;